import { ZodObject } from "zod";
import {Snippet} from "svelte";

export type SignProps = {
  defaultValues: any,
  signText: string,
  subSignText: string,
  subFormText: string,
  signButtonText: string,
  subFormLink: string,
  signButtonLinkText: string,
  submitAction: (data: any) => Promise<any>,
  googleButtonText: string,
  schema: ZodObject<any>,
  children?: Snippet,
  setLoading?: (loading: boolean) => void,
  showEmailSentMessage?: boolean,
  setShowEmailSentMessage?: (showEmailSentMessage: boolean) => void,
}
