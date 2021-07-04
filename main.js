const Discord = require('discord.js');
const client = new Discord.Client();
var fs = require('fs');



client.once('ready', () => {
	console.log('Ready!');
});


client.on('message', message => {
	if(message.content.startsWith(`https://linkvertise.com/`)){
        try {
            let first = message.content.split('=')
            let second = first[1].split('%')
            let final = second[0]
    
            var buf = Buffer.from(final, 'base64'); // Ta-da
    
    
    
    
            message.channel.send('<@'+message.author.id+'>'+'    '+ buf.toString());
        } catch (error) {
            message.channel.send('<@'+message.author.id+'>'+' We encountered an error decoding this URL');
        }

	}
});



client.login('');
