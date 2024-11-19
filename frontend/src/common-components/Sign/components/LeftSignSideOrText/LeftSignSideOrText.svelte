<script lang="ts">
    import DividerElement from "@/common-components/DividerElement/DividerElement.svelte";
    import LeftSignSideOrTextLayout
        from "@/common-components/Sign/components/LeftSignSideOrText/LeftSignSideOrTextLayout.svelte";
</script>

<LeftSignSideOrTextLayout>
    <DividerElement/>
    <div class="absolute -top-3 font-medium text-base left-0 text-center w-full">
        <span class={"bg-light-palette-background-block dark:bg-dark-palette-background-block px-2.5"}>or</span>
    </div>
</LeftSignSideOrTextLayout>