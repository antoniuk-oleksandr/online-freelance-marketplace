<script lang="ts">
    import LeftSignSideTopTextLayout
        from "@/common-components/Sign/components/LeftSignSideTopText/LeftSignSideTopTextLayout.svelte";

    type LeftSignSideTopTextProps = {
        signText: string,
        subSignText: string
    }

    const {signText, subSignText}: LeftSignSideTopTextProps = $props();
</script>

<LeftSignSideTopTextLayout>
    <h2 class="font-bold text-4xl">{signText}</h2>
    <p
            class="text-lg text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
    >{subSignText}</p>
</LeftSignSideTopTextLayout>
