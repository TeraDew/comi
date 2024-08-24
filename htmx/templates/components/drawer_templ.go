// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/yumenaka/comi/htmx/state"

func Drawer(s *state.GlobalState, slot templ.Component) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- drawer component --><!-- https://flowbite.com/docs/components/drawer/ --><div id=\"drawer-right\" x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(serverHostBindStr(s.ServerStatus.ServerHost))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/drawer.templ`, Line: 8, Col: 77}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"fixed top-0 right-0 z-40 w-64 h-screen p-4 overflow-y-auto transition-transform translate-x-full bg-base-100 text-base-content\" tabindex=\"-1\" aria-labelledby=\"drawer-right-label\"><h5 x-text=\"i18next.t(&#39;ReaderSettings&#39;)\" id=\"drawer-right-label\" class=\"inline-flex items-center mb-4 text-base font-semibold text-gray-500 dark:text-gray-400\"><svg class=\"w-4 h-4 me-2.5\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\"><path d=\"M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z\"></path></svg>Right drawer</h5><button type=\"button\" data-drawer-hide=\"drawer-right\" aria-controls=\"drawer-right\" class=\"font-bold bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 absolute top-2.5 end-2.5 inline-flex items-center justify-center dark:hover:bg-gray-600 dark:hover:text-white\">X</button><div class=\"flex flex-col items-center justify-center p-1 my-2 dbg-accent text-accent-content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = slot.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!--<div class=\"w-full font-semibold text-left\">主题</div> --><select x-model=\"theme\" x-on:change=\"theme = $event.target.value;console.log(theme);\" class=\"w-full h-10 m-2 border rounded bg-base-100 text-accent-content focus:outline-none\"><option value=\"retro\">Retro</option> <option value=\"light\">Light</option> <option value=\"dark\">Dark</option> <option value=\"dracula\">Dracula</option> <option value=\"cupcake\">Cupcake</option> <option value=\"cyberpunk\">Cyberpunk</option> <option value=\"valentine\">Valentine</option> <option value=\"aqua\">Aqua</option> <option value=\"lofi\">Lofi</option> <option value=\"halloween\">Halloween</option> <option value=\"coffee\">Coffee</option> <option value=\"winter\">Winter</option> <option value=\"nord\">Nord</option></select><!--选择语言的select --><!-- <div class=\"w-full font-semibold text-left\">语言</div> --><select x-model=\"i18next.language\" x-on:change=\"i18next.changeLanguage($event.target.value).then(location.reload())\" class=\"w-full h-10 pl-3 pr-3 m-2 border rounded bg-base-100 text-accent-content focus:outline-none\"><option value=\"en\">English</option> <option value=\"zh\">中文</option> <option value=\"ja\">日本語</option></select></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
