# bettelegram_tools

<h3>what is it</h3>
<p>Some tools from my telegram cli that i find decent enough to post publicly</p>

<h3>implemented so far:</h3>

<ul>
    <li>
        <div>
            <p>text to speech generated voice messages</p>
        </div>
    </li>
    <li>
        <div>
            <p>sending audio files as a voice note</p>
        </div>
    </li>
</ul>

<h3>use examples:</h3>
<p>sending a taunting voice message followed by a deepfried mp3 file</p>
<video src = "https://user-images.githubusercontent.com/55796857/187624503-73a927a1-a5cf-4b8a-ae05-cd74c665fbe0.mp4" width=180 ></video>


<h3>TODO:</h3>

<ul>
    <li> <p> sending any video as a videoNote </p>
    <li> <p> custom waveforms for a voiceNote </p>
    <li> <p> move tts language and chat list limit to command line args </p>
</ul>

<h3>manual:</h3>

```txt

send a regular text message
?: any text that doesnt start with a '\'

commands:

\v some text - send a tts voice note
\s path - send audio file as a voice note

\q select another chat
\e exit the program

```

