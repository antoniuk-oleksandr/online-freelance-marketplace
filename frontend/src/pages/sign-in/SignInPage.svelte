<script lang="ts">
    import {initialSignInData, signInSchema} from "@/pages/sign-in/helpers.ts";
    import SignInInputs from "@/pages/sign-in/components/SignInInputs/SignInInputs.svelte";
    import {handleSignIn} from "@/pages/sign-in/handlers.ts";
    import Sign from "@/common-components/Sign/Sign.svelte";

</script>

<Sign
        defaultValues={initialSignInData}
        googleButtonText={"Sign in with Google"}
        submitAction={handleSignIn}
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