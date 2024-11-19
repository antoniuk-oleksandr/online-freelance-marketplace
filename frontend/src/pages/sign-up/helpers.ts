import {z} from 'zod';
import {SignUpData} from "@/types/SignData.ts";

export const signUpSchema = z.object({
    username: z.string().min(1, {message: "Username is required."}),
    firstName: z.string().min(1, {message: "First name is required."}),
    surname: z.string().min(1, {message: "Surname is required."}),
    password: z.string().min(1, {message: "Password is required."}),
})

export const initialSignUpData: SignUpData = {
    firstName: "",
    surname: "",
    username: "",
    password: "",
}