import Sign from "@/common-components/Sign/Sign";
import {handleSignUp} from "@/pages/sign-up/handlers";
import {defaultSignUpValues, signUpSchema} from "@/pages/sign-up/helpers";
import SignUpInputs from "@/pages/sign-up/components/SignUpInputs/SignUpInputs";

const SignUpPage = () => {
    return (
        <Sign
            defaultValues={defaultSignUpValues}
            googleButtonText={"Continue with Google"}
            onSubmit={(data: any) => handleSignUp(data)}
            subSignText={"Sign up to enjoy all the features."}
            signText={"Sign Up"}
            subFormText={"Already have an account?"}
            signButtonText={"Sign Up"}
            subFormLink={"/sign-in"}
            signButtonLinkText={"Sign In"}
            schema={signUpSchema}
        >
            <SignUpInputs/>
        </Sign>
    )
}

export default SignUpPage;