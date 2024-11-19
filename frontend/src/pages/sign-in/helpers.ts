import {z} from 'zod';
import {SignInData} from "@/types/SignData.ts";

export const signInSchema = z.object({
    usernameOrEmail: z.string().min(1, {message: "Username or Email is required."}),
    password: z.string()
        .min(1, {message: "Password is required."})
        .min(8, {message: "Password must be at least 8 characters long."})
        .regex(/[a-z]/, { message: "Password must contain at least one lowercase letter." })
        .regex(/[A-Z]/, { message: "Password must contain at least one uppercase letter." })
        .regex(/[0-9]/, { message: "Password must contain at least one digit." })
        .regex(/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+/, { message: "Password must contain at least one special character." }),
})

export const initialSignInData: SignInData = {
    usernameOrEmail: "",
    password: "",
}