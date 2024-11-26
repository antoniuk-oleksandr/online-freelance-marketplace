<script lang="ts">
    import {handleConfirmEmailResponse} from "@/pages/confirm-email/helpers.ts";
    import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
    import Spinner from "@/common-components/Spinner/Spinner.svelte";
    import ConfirmEmailPageLayout from "@/pages/confirm-email/ConfirmEmailPageLayout.svelte";
    import ErrorComponent from "@/common-components/ErrorComponent/ErrorComponent.svelte";
    import {postAuthRequest} from "@/api/post-auth-request.ts";
    import SuccessMessage from "@/common-components/SuccessMessage/SuccessMessage.svelte";

    let error = $state<undefined | null | ResponseErrorEnum>(undefined);
    const setError = (value: undefined | null | ResponseErrorEnum) => error = value;

    $effect(() => {
        const params = new URLSearchParams(window.location.search);
        const token = params.get('token');

        if (!token) {
            setError(ResponseErrorEnum.InvalidToken);
            return;
        }
        postAuthRequest("confirm-email", token).then((response) => {
            handleConfirmEmailResponse(response, setError);
        });
    })
</script>

<ConfirmEmailPageLayout>
    {#if error === undefined}
        <Spinner/>
    {:else if error === null}
        <SuccessMessage
                title="Email Confirmed"
                description="Your email has been successfully confirmed. You can now proceed to use your account."
                subText="Thank you for verifying your email. We’re excited to have you on board!"
        />
    {:else if error !== null}
        <ErrorComponent error={error}/>
    {/if}
</ConfirmEmailPageLayout>
