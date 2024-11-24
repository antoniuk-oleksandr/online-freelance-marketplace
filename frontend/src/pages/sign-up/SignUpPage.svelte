<script lang="ts">
    import Sign from "@/common-components/Sign/Sign.svelte";
    import {initialSignUpData, signUpSchema} from "@/pages/sign-up/helpers.ts";
    import SignUpInputs from "@/pages/sign-up/components/SignUpInputs/SignUpInputs.svelte";
    import {postSignUpRequest} from "@/api/post-sign-up-request.ts";

    let showEmailSentMessage = $state(false);
    const setShowEmailSentMessage = (value: boolean) => showEmailSentMessage = value;
</script>

<Sign
        showEmailSentMessage={showEmailSentMessage}
        setShowEmailSentMessage={setShowEmailSentMessage}
        defaultValues={initialSignUpData}
        googleButtonText={"Continue with Google"}
        submitAction={postSignUpRequest}
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