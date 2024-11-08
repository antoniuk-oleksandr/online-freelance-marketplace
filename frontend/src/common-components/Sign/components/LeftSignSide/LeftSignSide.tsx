import LeftSignSideLayout from "./LeftSignSideLayout";
import {PasswordInput} from "@mantine/core";
import MyButton from "@/common-components/MyButton/MyButton";
import LeftSignSideTopText from "@/common-components/Sign/components/LeftSignSideTopText/LeftSignSideTopText";
import MyInput from "@/common-components/MyInput/MyInput";
import LeftSignSideOrText from "@/common-components/Sign/components/LeftSignSideOrText/LeftSignSideOrText";
import Link from "next/link";
import GoogleButton from "@/common-components/Sign/components/GoogleButton/GoogleButton";
import SubFormText from "@/common-components/Sign/components/SubFormText/SubFormText";
import SignForm from "@/common-components/Sign/components/SignForm/SignForm";
import {SignProps} from "@/types/SignProps";

const LeftSignSide = (props: SignProps) => {
    return (
        <LeftSignSideLayout>
            <LeftSignSideTopText {...props}/>
            <SignForm {...props}/>
            <GoogleButton {...props}/>
        </LeftSignSideLayout>
    )
}

export default LeftSignSide;