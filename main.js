const Discord = require('discord.js');
const client = new Discord.Client();

var fs = require('fs');



client.once('ready', () => {
	console.log('Ready!');
});


client.on('message', message => {
	if(message.content.startsWith(`https://quizlet.com/`)){

	}
});



client.login('');
