<script lang="ts">
    import {initialSignInData, signInSchema} from "@/pages/sign-in/helpers.ts";
    import SignInInputs from "@/pages/sign-in/components/SignInInputs/SignInInputs.svelte";
    import Sign from "@/common-components/Sign/Sign.svelte";
    import {postSignInRequest} from "@/api/post-sign-in-request.ts";

</script>

<Sign
        defaultValues={initialSignInData}
        googleButtonText={"Sign in with Google"}
        submitAction={postSignInRequest}
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