<script lang="ts">
    import {postConfirmEmailRequest} from "@/api/post-confirm-email-request.ts";
    import {handleConfirmEmailResponse} from "@/pages/confirm-email/helpers.ts";
    import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
    import Spinner from "@/common-components/Spinner/Spinner.svelte";
    import ConfirmEmailPageLayout from "@/pages/confirm-email/ConfirmEmailPageLayout.svelte";
    import ErrorComponent from "@/common-components/ErrorComponent/ErrorComponent.svelte";
    import EmailConfimedBlock from "@/pages/confirm-email/components/EmailConfimedBlock/EmailConfimedBlock.svelte";

    let error = $state<undefined | null | ResponseErrorEnum>(undefined);
    const setError = (value: undefined | null | ResponseErrorEnum) => error = value;

    $effect(() => {
        const params = new URLSearchParams(window.location.search);
        const token = params.get('token');

        if (!token) {
            setError(ResponseErrorEnum.InvalidToken);
            return;
        }
        postConfirmEmailRequest(token).then((response) => {
            handleConfirmEmailResponse(response, setError);
        });
    })
</script>

<ConfirmEmailPageLayout>
    {#if error === undefined}
        <Spinner/>
    {:else if error === null}
        <EmailConfimedBlock/>
    {:else if error !== null}
        <ErrorComponent error={error}/>
    {/if}
</ConfirmEmailPageLayout>
