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

func Login_Page() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n\t// page script starts\n\tconst handleLoginFormSubmit = () => {\n\t\tconst loginButton = document.getElementById(\"loginButton\")\n\t\tloginButton.innerHTML = `<span class=\"loading loading-spinner loading-lg\"></span>`\n\t}\n\tdocument.body.addEventListener(\"htmx:afterRequest\", (event) => {\n\t\tconst loginButton = document.getElementById(\"loginButton\")\n\t\tloginButton.innerHTML = \"Login\"\n\t\tconst res = JSON.parse(event.detail.xhr.response)\n\t\tif (res.status_code != \"PASS_LGN\") {\n\t\t\tdocument.body.insertAdjacentHTML('beforeend', res.toast);\n\t\t\tsetTimeout(() => {\n\t\t\t\tdocument.getElementById(\"loginToast\").remove();\n\t\t\t}, 2000);\n\t\t}\n\t})\n\t// page script ends\n</script> <div class=\"flex flex-col items-center mt-[150px] gap-[20px]\"><div class=\"flex flex-col items-center gap-[15px]\"><h1 class=\"text-4xl\">Login</h1><p>Don't have an account? <a href=\"/register\" class=\"link link-primary\">Register</a></p></div><form hx-post=\"/auth/login\" hx-swap=\"none\" name=\"loginForm\" class=\"flex flex-col items-center gap-[15px] w-full\"><input name=\"username\" placeholder=\"Username\" autocomplete=\"off\" class=\"input border-neutral focus:outline-none focus:border-primary placeholder-base-200 w-full max-w-[250px]\"> <input name=\"password\" placeholder=\"Password\" autocomplete=\"off\" class=\"input border-neutral focus:outline-none focus:border-primary placeholder-base-200 w-full max-w-[250px]\"> <button hx-on:click=\"handleLoginFormSubmit()\" id=\"loginButton\" type=\"submit\" class=\"btn btn-primary w-full max-w-[180px]\">Login</button></form></div>")
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
