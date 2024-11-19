import SignFormLayout from "./SignFormLayout";
import MyButton from "@/common-components/MyButton/MyButton";
import LeftSignSideOrText from "@/common-components/Sign/components/LeftSignSideOrText/LeftSignSideOrText";
import SubFormText from "@/common-components/Sign/components/SubFormText/SubFormText";
import {SignProps} from "@/types/SignProps";

const SignForm = (props: SignProps) => {
    const {children, signButtonText} = props;

    return (
        <SignFormLayout {...props}>
            {children}
            <MyButton
                type={"submit"}
            >{signButtonText}</MyButton>
            <LeftSignSideOrText/>
            <SubFormText {...props}/>
        </SignFormLayout>
    )
}

export default SignForm;