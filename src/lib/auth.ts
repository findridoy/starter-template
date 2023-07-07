import { ofetch } from "ofetch"
import { redirect } from "react-router-dom"


export const isLoggedIn = async () => {
  try {
    await ofetch("/api/v1/me", {baseURL: "http://localhost:8080", credentials: "include"})
    return true  
  } catch (err) {
    
  }
  return false 
}
