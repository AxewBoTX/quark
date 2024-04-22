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

func Register_Page() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n\t// page script ends\n\tconst handleRegisterFormSubmit = () => {\n\t\tconst registerButton = document.getElementById(\"registerButton\")\n\t\tregisterButton.innerHTML = `<span class=\"loading loading-spinner loading-lg\"></span>`\n\t}\n\tdocument.body.addEventListener(\"htmx:afterRequest\", (event) => {\n\t\tconst registerButton = document.getElementById(\"registerButton\")\n\t\tregisterButton.innerHTML = \"Register\"\n\t\tconst res = JSON.parse(event.detail.xhr.response)\n\t\tdocument.body.insertAdjacentHTML('beforeend', res.toast);\n\t\tsetTimeout(() => {\n\t\t\tdocument.getElementById(\"registerToast\").remove();\n\t\t}, 2000);\n\t})\n\t// page script ends\n</script> <div class=\"flex flex-col items-center mt-[50px] gap-[20px]\"><h1 class=\"text-4xl\">Register</h1><form hx-post=\"/auth/register\" hx-swap=\"none\" hx-trigger=\"submit\" name=\"registerForm\" class=\"flex flex-col items-center gap-[15px] w-full\"><input name=\"username\" placeholder=\"Username\" autocomplete=\"off\" class=\"input input-bordered w-full max-w-[250px]\"> <input name=\"password\" placeholder=\"Password\" autocomplete=\"off\" class=\"input input-bordered w-full max-w-[250px]\"> <input name=\"passwordConfirm\" placeholder=\"Confirm Password\" autocomplete=\"off\" class=\"input input-bordered w-full max-w-[250px]\"> <button hx-on:click=\"handleRegisterFormSubmit()\" id=\"registerButton\" type=\"submit\" class=\"btn w-full max-w-[180px]\">Register</button></form></div>")
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
