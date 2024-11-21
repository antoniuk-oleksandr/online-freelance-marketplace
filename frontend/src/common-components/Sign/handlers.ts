import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
import {ResponseMessageEnum} from "@/types/ResponseMessageEnum.ts";

export const handleGoogleSignInButtonClick = (
    tokenResponse: any
) => {
    console.log(tokenResponse)
}

export const handleSignSubmit = async (
    values: any,
    submitAction: (values: any) => Promise<any>,
    setErrors: (fieldName: string, error: string) => void,
    setLoading?: (loading: boolean) => void,
    setShowEmailSentMessage?: (showEmailSentMessage: boolean) => void
) => {
    if (!setLoading) return;
    setLoading(true);

    const {data, status} = await submitAction(values);

    if (status === 200 && data.message && data.message === ResponseMessageEnum.EmailSentSuccessfully) {
        setShowEmailSentMessage && setShowEmailSentMessage(true);
    } else if (data.error === ResponseErrorEnum.UserAlreadyExists) {
        setErrors("username", "User already exists.")
    }

    setLoading(false);
}