import Sign from "@/common-components/Sign/Sign";
import {handleSignIn} from "@/pages/sign-in/handlers";
import {defaultSignInValues, signInSchema} from "@/pages/sign-in/helpers";
import {useGoogleLogin} from "@react-oauth/google";
import SignInInputs from "@/pages/sign-in/components/SignInInputs/SignInInputs";

const SignInPage = () => {
    return (
        <Sign
            defaultValues={defaultSignInValues}
            googleButtonText={"Sign in with Google"}
            onSubmit={handleSignIn}
            signText={"Sign In"}
            subSignText={"Please sign in to continue to your account."}
            subFormText={"Don't have an account?"}
            signButtonText={"Sign In"}
            subFormLink={"/sign-up"}
            signButtonLinkText={"Sign Up"}
            schema={signInSchema}
        >
            <SignInInputs/>
        </Sign>
    )
}

export default SignInPage;