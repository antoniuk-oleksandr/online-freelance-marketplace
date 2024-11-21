<script lang="ts">
    import type {SignProps} from "@/types/SignProps.ts";
    import SignLayout from "@/common-components/Sign/SignLayout.svelte";
    import LeftSignSide from "@/common-components/Sign/components/LeftSignSide/LeftSignSide.svelte";
    import RightSignSide from "@/common-components/Sign/components/RightSignSide/RightSignSide.svelte";
    import SentEmailMessage from "@/common-components/Sign/components/SentEmailMessage/SentEmailMessage.svelte";

    const signProps: SignProps = $props();
</script>

{#if signProps.showEmailSentMessage}
    <SentEmailMessage/>
{:else}
    <SignLayout>
        <LeftSignSide {...signProps}/>
        <RightSignSide/>
    </SignLayout>
{/if}
