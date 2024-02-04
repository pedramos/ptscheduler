// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type Login struct {
	header templ.Component
	body   templ.Component
}

func NewLogin() Login {
	return Login{
		header: loginHeader(),
		body:   loginBody(),
	}
}

func (p Login) Header() templ.Component { return p.header }
func (p Login) Body() templ.Component   { return p.body }

func loginHeader() templ.Component {
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
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func loginBody() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"bg-gray-50 dark:bg-gray-900\"><div class=\"flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0\"><a href=\"#\" class=\"flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white\"><img class=\"w-8 h-8 mr-2\" src=\"/static/training-gym.svg\" alt=\"logo\"> Há Cadinho para todos</a><div class=\"w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700\"><div class=\"p-6 space-y-4 md:space-y-6 sm:p-8\"><h1 class=\"text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white\">Sign in to your account</h1><form class=\"space-y-4 md:space-y-6\" action=\"/login\" method=\"post\"><div><label for=\"email\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Your email</label> <input type=\"username\" name=\"username\" id=\"username\" class=\"bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" placeholder=\"username\" required=\"\"></div><div><label for=\"password\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Password</label> <input type=\"password\" name=\"password\" id=\"password\" placeholder=\"••••••••\" class=\"bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" required=\"\"></div><div class=\"flex items-center justify-between\"><div class=\"flex items-start\"><!--\n\t\t\t\t\t\t\t\t<div class=\"flex items-center h-5\">\n\t\t\t\t\t\t\t\t\t<input id=\"remember\" aria-describedby=\"remember\" type=\"checkbox\" class=\"w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-primary-600 dark:ring-offset-gray-800\"/>\n\t\t\t\t\t\t\t\t</div>\n\n\t\t\t\t\t\t\t\t<div class=\"ml-3 text-sm\">\n\t\t\t\t\t\t\t\t\t<label for=\"remember\" class=\"text-gray-500 dark:text-gray-300\">Remember me</label>\n\t\t\t\t\t\t\t\t</div>\n--></div><a href=\"#\" class=\"text-sm font-medium text-primary-600 hover:underline dark:text-primary-500\">Forgot password?</a></div><button type=\"submit\" class=\"w-full text-white bg-blue-600 hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800\">Sign in</button><!--\n\t\t\t\t\t\t<p class=\"text-sm font-light text-gray-500 dark:text-gray-400\">\n\t\t\t\t\t\t\tDon’t have an account yet? <a href=\"#\" class=\"font-medium text-primary-600 hover:underline dark:text-primary-500\">Sign up</a>\n\t\t\t\t\t\t</p>\n--></form></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
