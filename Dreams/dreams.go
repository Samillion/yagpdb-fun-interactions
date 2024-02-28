{{ $ch := 1200964297518497903 }}
{{ $minimum := 20 }}
{{ $cd := 600 }}

{{ $cmd := .Cmd }}
{{ $msg := trimSpace .StrippedMsg }}

{{ if not (getChannel $ch) }}
	{{ sendMessage nil (print "**Error:** Please edit the command code and adjust `$ch` to a valid channel ID.") }}
	{{ deleteTrigger 0 }}
	{{ return }}
{{ end }}

{{ if $msg }}
	{{ if gt (len $msg) $minimum }}
		{{ $db := "dreams" }}
		{{ $ccid := .CCID }}
		{{ if not (dbGet $ccid $db) }}
			{{ $note := print "Use " $cmd " to post your own dream." }}
			{{ $embed := cembed 
				"title" "Dream"
				"description" $msg 
				"color" 8421888
				"footer" (sdict "text" $note)
			}}
			{{ sendMessage $ch $embed }}
			{{ dbSetExpire $ccid $db "cooldown" $cd }}
			{{ $cdText := print 
				"A global cooldown is active for 10 minutes, meaning no one can post a dream during that time." "\n\n"
				"__This messaage will be deleted automatically when the cooldown is over__."
			}}
			{{ $cdEmbed := cembed 
				"title" "Cooldown"
				"description" $cdText
				"color" 11993101
			}}
			{{ $x := sendMessageRetID $ch $cdEmbed }}
			{{ deleteMessage $ch $x $cd }}
		{{ else }}
			{{ $x := sendMessageRetID nil "There is a global 10 minutes cooldown per dream." }}
			{{ deleteMessage nil $x 10 }}
		{{ end }}
	{{ else }}
		{{ $x := sendMessageRetID nil (print "Your dream post must contain at least " $minimum " characters.") }}
		{{ deleteMessage nil $x 10 }}
	{{ end }}
	{{ deleteTrigger 0 }}
{{ else }}
	{{ $usage := print "**Usage:**" "\n" "```" $cmd " your message```" }}
	{{ $explain := print 
		"Had an interesting dream last night? Share it with us by using this command, it will post your message in <#" $ch ">." "\n\n" 
		$usage "\n" 
		"**Explanation:**" "\n"
		"- Once you use `" $cmd "` it will post your message in the relative channel." "\n"
		"- Your message will be deleted instantly, leaving only the one in <#" $ch ">." "\n"
		"- If the dream is less than " $minimum " characters, nothing will be posted." "\n"
		"- A limit of one dream every 10 minutes, globally." "\n"
		"- Feel free to review the code of this command [here](https://github.com/Samillion/yagpdb-cc/tree/main/Dreams)."
	}}
	{{ $main := cembed 
		"title" "Dreams"
		"description" $explain
	}}
	{{ sendMessage nil $main }}
{{ end }}
