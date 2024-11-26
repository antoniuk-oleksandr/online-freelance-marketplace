<script lang="ts">
    import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
    import {processUrlToken, initialResetPasswordData, resetPasswordSchema} from "@/pages/reset-password/helpers.ts";
    import ResetPasswordInputs from "@/pages/reset-password/components/ResetPasswordInputs.svelte";
    import Sign from "@/common-components/Sign/Sign.svelte";
    import ResetPasswordPageLayout from "@/pages/reset-password/ResetPasswordPageLayout.svelte";
    import ErrorComponent from "@/common-components/ErrorComponent/ErrorComponent.svelte";
    import {handlePasswordResetRequest} from "@/pages/reset-password/handlers.ts";
    import SuccessMessage from "@/common-components/SuccessMessage/SuccessMessage.svelte";

    let token = $state<string>("");
    const setToken = (value: string) => token = value;

    let error = $state<undefined | null | ResponseErrorEnum>(undefined);
    const setError = (value: undefined | null | ResponseErrorEnum) => error = value;

    let reset = $state<boolean | undefined>();
    const setReset = (value: boolean) => reset = value;

    $effect(() => processUrlToken(setToken, setError))

    const postSignInRequest = (body: any) =>
        handlePasswordResetRequest(body, token, setError, setReset);
</script>

<ResetPasswordPageLayout>
    {#if reset}
        <SuccessMessage
            title="Password Reset"
            description="Your password has been successfully reset. You can now sign in with your new password."
            subText="Thank you for taking the time to secure your account. We're here to help if you need further assistance."
        />
    {:else if error === undefined}
        <Sign
                defaultValues={initialResetPasswordData}
                submitAction={postSignInRequest}
                signText={"Reset Password"}
                subSignText={"Please enter your new password"}
                signButtonText={"Reset Password"}
                subFormLink={"/sign-up"}
                signButtonLinkText={"Sign Up"}
                schema={resetPasswordSchema}
        >
            <ResetPasswordInputs/>
        </Sign>
    {:else if error !== null}
        <ErrorComponent error={error}/>
    {/if}
</ResetPasswordPageLayout>
