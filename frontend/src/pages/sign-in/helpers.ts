import { z } from 'zod';
import { SignInData } from "@/types/SignData";

export const signInSchema = z.object({
    usernameOrEmail: z.string().min(1, { message: "Username or Email is required." }),
    password: z.string().min(1, { message: "Password is required." }),
})

export const initialSignInData: SignInData = {
    usernameOrEmail: "",
    password: "",
}
