<script lang="ts">
    import SubFormTextLayout from "@/common-components/Sign/components/SubFormText/SubFormTextLayout.svelte";
    import {Link} from "svelte-routing";

    type SubFormTextProps = {
        subFormText: string | undefined,
        signButtonLinkText: string | undefined,
        subFormLink: string | undefined,
    }

    const {subFormText, signButtonLinkText, subFormLink}: SubFormTextProps = $props();
</script>


{#if subFormText !== undefined && signButtonLinkText !== undefined && subFormLink !== undefined}
    <SubFormTextLayout>
        <span>{subFormText}</span>
        <Link
                class="text-cyan-500 hover:underline ease-out duration-200"
                to={subFormLink}
        >{signButtonLinkText}</Link>
    </SubFormTextLayout>
{/if}
