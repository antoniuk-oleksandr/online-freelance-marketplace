<script lang="ts">
    import {initialSignInData, signInSchema} from "@/pages/sign-in/helpers.ts";
    import SignInInputs from "@/pages/sign-in/components/SignInInputs/SignInInputs.svelte";
    import Sign from "@/common-components/Sign/Sign.svelte";
    import {postAuthRequest} from "@/api/post-auth-request.ts";
    import type {SignInData} from "@/types/SignData.ts";

    const postSignInRequest = async (body: SignInData) =>
        postAuthRequest("sign-in", undefined, body)
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