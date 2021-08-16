var messageInput = document.getElementById('messageInput')
var state // channel id
var refresh // reload messages

class Message {
  constructor (go) {
    this.ID = go.id
    this.Content = go.content
    this.Author = go.author
    this.Element = this.Div()
  }

  Div () {
    var e = document.createElement('div')
    e.classList.add('message')
    e.innerHTML = `${this.Author.username}:${this.Content}`
    return e
  }
}

class Channel {
  constructor (go) {
    this.ID = go.id
    this.Name = go.name
    this.Type = go.type
    this.Element = this.Div()
  }

  Div () {
    var e = document.createElement('div')
    e.classList.add('channel')
    e.innerHTML = this.Name
    e.onclick = this.loadMessages.bind(this)
    return e
  }

  async loadMessages () {
    state = this.ID
    refresh = this.loadMessages.bind(this)
    document.getElementById('messages').innerHTML = ''
    var messages = await messagesGO(this.ID)
    for (let m of messages) {
      var message = new Message(m)
      document.getElementById('messages').append(message.Element)
    }
  }
}
class Guild {
  constructor (go) {
    this.ID = go.id
    this.Name = go.name
    this.Icon = `https://cdn.discordapp.com/icons/${this.ID}/${go.icon}.png`
    this.Element = this.image()
  }

  image () {
    var e = document.createElement('img')
    e.classList.add('icon')
    e.src = this.Icon
    e.onclick = this.loadChannels.bind(this)
    return e
  }

  async loadChannels () {
    document.getElementById('channels').innerHTML = ''
    var channels = await channelsGO(this.ID)
    for (let c of channels) {
      var channel = new Channel(c)
      if (c.type == 'topic') continue
      document.getElementById('channels').append(channel.Element)
    }
  }
}

document.onkeypress = function (e) {
  if (e.key == 'Enter') {
    sendGO(state, messageInput.value)
    refresh()
  }
}
;(async () => {
  var guilds = await guildsGO()
  for (const g of guilds) {
    var guild = new Guild(g)
    document.getElementById('guilds').append(guild.Element)
  }
  setInterval(refresh, 5000)
})()
