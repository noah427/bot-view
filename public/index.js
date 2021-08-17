var messageInput = document.getElementById('messageInput')
var state // channel id
var refresh // reload messages

class Voice {
  constructor (channel) {
    this.Element = this.Div()
  }

  Div () {
    var e = document.createElement('div')
    e.classList.add('voice')
    e.append(this.DisconnectButton())
    return e
  }

  DisconnectButton () {
    var e = document.createElement('div')
    e.classList.add('disconnect')
    e.innerHTML = 'Disconnect'
    e.onclick = this.Disconnect.bind(this)
    return e
  }

  async Disconnect () {
    await disconnectVoiceGO()
    document.getElementById('messages').innerHTML = ''
  }
}

class Message {
  constructor (go) {
    this.ID = go.id
    this.Content = go.content
    this.Author = go.author
    if (go.attachments.length > 0) {
      this.attachment = go.attachments[0].url
    } else {
      this.attachment = null
    }
    this.Element = this.Div()
  }

  Pfp () {
    var e = document.createElement('img')
    e.classList.add('pfp')
    e.src = `https://cdn.discordapp.com/avatars/${this.Author.id}/${this.Author.avatar}.webp?size=128`
    return e
  }

  Image () {
    var e = document.createElement('img')
    e.classList.add('attachment')
    e.src = this.attachment

    return e
  }

  Username () {
    var e = document.createElement('p')
    e.innerHTML = `${this.Author.username} : `
    return e
  }

  Text () {
    var e = document.createElement('p')
    e.innerHTML = this.Content
    return e
  }

  UserInfo () {
    var e = document.createElement('div')
    e.classList.add('userinfo')
    e.append(this.Pfp())
    e.append(this.Username())
    return e
  }

  Div () {
    var e = document.createElement('div')
    e.classList.add('message')
    e.append(this.UserInfo())
    e.append(this.Text())
    if (this.attachment != null) {
      e.append(this.Image())
    }
    return e
  }
}

class Channel {
  constructor (go) {
    this.ID = go.id
    this.Name = go.name
    this.Type = go.type
    this.Guild = go.guild_id
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
    e.onclick =
      this.Type == 0 ? this.loadMessages.bind(this) : this.loadVoice.bind(this)
    e.append(this.Marker())
    e.append(this.Text())
    return e
  }

  async loadVoice () {
    console.log(this.Guild, this.ID)
    await joinVoiceGO(this.Guild, this.ID)
    document.getElementById('messages').innerHTML = ''
    var voice = new Voice(this)
    document.getElementById('messages').append(voice.Element)
  }

  async loadMessages () {
    state = this.ID
    refresh = this.loadMessages.bind(this)
    var messages = await messagesGO(this.ID)
    document.getElementById('messages').innerHTML = ''
    for (let m of messages) {
      var message = new Message(m)
      document.getElementById('messages').append(message.Element)
    }
    document.getElementById('messages').scrollBy(0, 10000)
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
  setInterval(callRefresh, 5000)
})()

function callRefresh () {
  state ? refresh() : undefined
}
