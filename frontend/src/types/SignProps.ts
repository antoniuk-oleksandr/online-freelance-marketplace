import {ZodType} from "zod";
import {Snippet} from "svelte";

export type SignProps = {
  defaultValues: any,
  submitAction: (data: any) => Promise<any>,
  signText: string,
  subSignText: string,
  schema: ZodType<any, any>,
  signButtonText: string,
  showBackButton?: boolean,
  signButtonLinkText?: string,
  subFormLink?: string,
  subFormText?: string,
  googleButtonText?: string,
  children?: Snippet,
  setLoading?: (loading: boolean) => void,
  showEmailSentMessage?: boolean,
  setShowEmailSentMessage?: (showEmailSentMessage: boolean) => void,
}
