import { z } from 'zod';
import { SignUpData } from "@/types/SignData";
import { generateECDHKeyPair } from '@/utils/ecdh-utils';
import { postAuthRequest } from '@/api/post-auth-request';
import { encryptWithKey, generateAESIV } from '@/utils/aes-utils';
import { encryptUserCredentials } from '@/common-components/Sign/helpers';

export const signUpSchema = z.object({
    username: z.string()
        .min(1, { message: "Username is required." })
        .min(3, { message: "Username must be at least 3 characters long." })
        .max(16, { message: "Username must be at most 16 characters long." })
        .regex(/^[a-zA-Z0-9_]+$/, { message: "Username must contain only letters, digits, and underscores." }),
    firstName: z.string().min(1, { message: "First name is required." }),
    surname: z.string().min(1, { message: "Surname is required." }),
    email: z.string()
        .min(1, { message: "Email is required." })
        .email({ message: "Invalid email address." }),
    password: z.string()
        .min(1, { message: "Password is required." })
        .min(8, { message: "Password must be at least 8 characters long." })
        .regex(/[a-z]/, { message: "Password must contain at least one lowercase letter." })
        .regex(/[A-Z]/, { message: "Password must contain at least one uppercase letter." })
        .regex(/[0-9]/, { message: "Password must contain at least one digit." })
        .regex(/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+/, { message: "Password must contain at least one special character." }),
})

export const initialSignUpData: SignUpData = {
    username: "",
    firstName: "",
    surname: "",
    password: "",
    email: "",
    privateKey: new Uint8Array(),
    publicKey: new Uint8Array(),
    privateKeyIV: new Uint8Array(),
    privateKeySalt: new Uint8Array(),
}

export const postSignUpRequest = async (body: SignUpData) => {
    const data = await encryptUserCredentials(body, body.password);

    return postAuthRequest('sign-up', true, data)
}
