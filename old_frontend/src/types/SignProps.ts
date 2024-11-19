import {ReactNode} from "react";
import {ZodObject} from "zod";

export type SignProps = {
    defaultValues: any,
    signText: string,
    subSignText: string,
    subFormText: string,
    signButtonText: string,
    subFormLink: string,
    signButtonLinkText: string,
    onSubmit: (data: any) => void,
    googleButtonText: string,
    children?: ReactNode,
    schema: ZodObject<any>,
}