# My AI is an AI assistant developed using the OpenAI API

## Setup Requirements

- Go 
- OpenAI API Key
- Create ~/.my-ai.conf in this form

```json
{
	"apikey": "<YOUR OPENAI API KEY>",
	"model": "gpt-3.5-turbo",
	"name": "My Assistant",
	"instructions": "You are a general assistant"
}
```

You can set the model to gpt-4, gpt-3.5, fpt-3.5-turbo

## Usage

```sh
$ my-ai
```
When a prompt appears, ask your questions. You can expand on an answer but sending another prompt.

ex:
```sh
> What is 2+2?
hmmm...

2 + 2 equals 4.

> add four more
hmmm...
If you add four more to 2 + 2, the result is 8.

> how do I solve (3x + 55) = (x - 10)
hmmm...
The solution to the equation (3x + 55) = (x - 10) is x = -65/2 or -32.5.
```

# Other

- This has only been tested and run on a Mac I'm not sure the config load will work in Windows but would be an easy change.
- This is pretty basic for now and only supports text -> text
