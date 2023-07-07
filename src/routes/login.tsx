import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ofetch } from "ofetch";
import { Form, redirect } from "react-router-dom";

export default function Login() {
  return (
    <>
      <Form method="post" action="/login" className="grid w-full max-w-sm items-center gap-1.5 mx-auto min-h-screen content-center">
        <Label>Username</Label>
        <Input className="max-w-sm" required name="identity" />
        <Label>Password</Label>
        <Input className="max-w-sm" type="password" required name="password" />
        <Button type="submit">Login</Button>
      </Form>
    </>
  )
}

export async function LoginRouteAction({ request }) {
  const formData = await request.formData()



  if (request.method === "POST") {
    const { token, env } = await ofetch("/api/v1/login", { method: 'POST', baseURL: 'http://localhost:8080', body: formData }).catch(err => {
      return {success: false}
    })


    if (env.toLowerCase() === "development") {
      
      document.cookie = `Authorization=${token}`
    }

    return redirect("/")
  }

  throw new Response("", { status: 405 })
}
