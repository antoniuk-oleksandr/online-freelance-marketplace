import SubFormTextLayout from "./SubFormTextLayout";
import Link from "next/link";
import {SignProps} from "@/types/SignProps";

const SubFormText = (props: SignProps) => {
    const {subFormText, signButtonLinkText, subFormLink} = props;

    return (
        <SubFormTextLayout>
            <span>{subFormText}</span>
            <Link
                className={"text-cyan-500 hover:underline ease-out duration-200"}
                href={subFormLink}
            >{signButtonLinkText}</Link>
        </SubFormTextLayout>
    )
}

export default SubFormText;