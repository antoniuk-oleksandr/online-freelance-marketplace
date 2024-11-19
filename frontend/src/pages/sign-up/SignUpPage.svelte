<script lang="ts">
    import Sign from "@/common-components/Sign/Sign.svelte";
    import {initialSignUpData, signUpSchema} from "@/pages/sign-up/helpers.ts";
    import {handleSignUp} from "@/pages/sign-up/handlers.ts";
    import SignUpInputs from "@/pages/sign-up/components/SignUpInputs/SignUpInputs.svelte";

</script>

<Sign
        defaultValues={initialSignUpData}
        googleButtonText={"Sign in with Google"}
        submitAction={handleSignUp}
        signText={"Sign Up"}
        subSignText={"Please sign Up to continue to your account."}
        subFormText={"Already have an account?"}
        signButtonText={"Sign Up"}
        subFormLink={"/sign-in"}
        signButtonLinkText={"Sign In"}
        schema={signUpSchema}
>
    <SignUpInputs/>
</Sign>