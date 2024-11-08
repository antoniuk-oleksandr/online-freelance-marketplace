import {CredentialResponse} from "@react-oauth/google";
import {jwtDecode} from "jwt-decode";

export const handleSignUpButtonClick = (response: CredentialResponse) => {
    const {credential} = response;

    if(!credential) return;

    const decoded = jwtDecode(credential);
    console.log(decoded);
}


export const handleSignUp = (data: any) => {
    console.log(data);
}