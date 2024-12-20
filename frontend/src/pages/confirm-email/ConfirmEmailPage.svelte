<script lang="ts">
    import {handleConfirmEmailResponse} from "@/pages/confirm-email/helpers.ts";
    import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
    import Spinner from "@/common-components/Spinner/Spinner.svelte";
    import ConfirmEmailPageLayout from "@/pages/confirm-email/ConfirmEmailPageLayout.svelte";
    import {postAuthRequest} from "@/api/post-auth-request.ts";
    import SuccessMessage from "@/common-components/SuccessMessage/SuccessMessage.svelte";
    import {errorStore} from "@/common-stores/error-store.ts";

    $effect(() => {
        const params = new URLSearchParams(window.location.search);
        const token = params.get('token');

        if (!token) {
            errorStore.set({shown: true, error: ResponseErrorEnum.InvalidToken});
            return;
        }
        postAuthRequest("confirm-email", token).then((response) => {
            handleConfirmEmailResponse(response);
        });
    })
</script>

<ConfirmEmailPageLayout>
    <SuccessMessage
            title="Email Confirmed"
            description="Your email has been successfully confirmed. You can now proceed to use your account."
            subText="Thank you for verifying your email. We’re excited to have you on board!"
    />
</ConfirmEmailPageLayout>
