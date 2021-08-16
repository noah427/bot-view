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

  Pfp () {
    var e = document.createElement('img')
    e.src = `https://cdn.discordapp.com/avatars/${this.Author.id}/${this.Author.avatar}.webp?size=128`
    return e
  }

  Text () {
    var e = document.createElement('p')
    e.innerHTML = `${this.Author.username}:${this.Content}`
    return e
  }

  Div () {
    var e = document.createElement('div')
    e.classList.add('message')
    e.append(this.Pfp())
    e.append(this.Text())
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

  Marker () {
    var e = document.createElement('img')
    e.src = this.Type == 0 ? 'svg/hashtag.svg' : 'svg/speaker.svg'
    return e
  }

  Text () {
    var e = document.createElement('p')
    e.innerHTML = this.Name
    return e
  }

  Div () {
    var e = document.createElement('div')
    e.classList.add('channel')
    e.onclick = this.loadMessages.bind(this)
    e.append(this.Marker())
    e.append(this.Text())
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
    document.getElementById('messages').scrollBy(0, 1000)
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

document.onkeypress = async function (e) {
  if (e.key == 'Enter') {
    await sendGO(state, messageInput.value)
    messageInput.value = ''
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
