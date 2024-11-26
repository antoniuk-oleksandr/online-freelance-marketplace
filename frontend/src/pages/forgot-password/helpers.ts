import {z} from "zod";

export const initialForgotPasswordData: any = {
    usernameOrEmail: "",
}

export const forgotPasswordSchema = z.object({
    usernameOrEmail: z.string().min(1, {message: "Username or email is required."}),
})