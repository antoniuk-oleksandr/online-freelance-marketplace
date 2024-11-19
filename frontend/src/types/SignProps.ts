import { ZodObject } from "zod";
import {Snippet} from "svelte";

export type SignProps = {
  defaultValues: any;
  signText: string;
  subSignText: string;
  subFormText: string;
  signButtonText: string;
  subFormLink: string;
  signButtonLinkText: string;
  submitAction: (data: any) => void;
  googleButtonText: string;
  children?: Snippet;
  schema: ZodObject<any>;
};
