// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package routes

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"quark/client/web/components"
)

func Chat_Page() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n\tdocument.body.addEventListener(\"htmx:wsBeforeSend\", (event) => {\n\t\tevent.preventDefault()\n\t\tconst message = {...JSON.parse(event.detail.message)}\n\t\tconst socket = event.detail.socketWrapper\n\t\tsocket.send(JSON.stringify(message))\n\t})\n\tdocument.body.addEventListener(\"htmx:wsAfterMessage\", (event) => {\n\t\tconst messageList = document.getElementById(\"MessageList\")\n\t\tconst message = JSON.parse(event.detail.message)\n\t\tdocument.getElementById(\"MessageInput\").value = \"\"\n\t\tswitch (message.type) {\n\t\t\tcase \"JOIN\":\n\t\t\t\tmessageList.insertAdjacentHTML(\"beforeend\", `<div class=\"flex justify-center text-semibold\">\n\t\t\t\t\t${message.username + \" joined the server\"}\n\t\t\t\t</div>`)\n\t\t\t\tbreak;\n\t\t\tcase \"MSG\":\n\t\t\t\tmessageList.insertAdjacentHTML(\"beforeend\", `<div class=\"flex flex-col items-start\">\n\t\t\t\t<p class=\"text-lg font-semibold bg-secondary text-secondary-content w-fit pl-1 pr-2 rounded-tl-lg rounded-tr-lg\" >${message.username}</p>\n\t\t\t\t<p class=\"bg-accent text-accent-content rounded-tl-none rounded-bl-badge rounded-r-badge px-4 py-1 max-w-[400px]\">${message.body}</p>\n\t\t</div>`)\n\t\t\t\tbreak;\n\t\t\tcase \"LEAVE\":\n\t\t\t\tmessageList.insertAdjacentHTML(\"beforeend\", `<div class=\"flex justify-center text-semibold\">\n\t\t\t\t\t${message.username + \" left the server\"}\n\t\t\t\t</div>`)\n\t\t\t\tbreak;\n\t\t}\n\t})\n</script> <div hx-ext=\"ws\" ws-connect=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(ctx.Value("realtime_server_addr").(string))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/routes/chat.templ`, Line: 40, Col: 74}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"flex justify-center w-full mb-[120px]\"><div id=\"MessageList\" class=\"flex flex-col w-full gap-[25px] p-3 max-w-[1200px]\"></div></div><div class=\"fixed left-0 right-0 bottom-0 flex items-center justify-center\"><form ws-send class=\"p-3 flex items-center justify-center gap-3 w-full max-w-[800px]\"><input id=\"MessageInput\" name=\"body\" placeholder=\"Message\" autocomplete=\"off\" class=\"input input-bordered w-full\"> <button type=\"submit\" class=\"btn btn-primary btn-wide max-w-[100px]\">Send</button></form></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = components.Base_HTML().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
