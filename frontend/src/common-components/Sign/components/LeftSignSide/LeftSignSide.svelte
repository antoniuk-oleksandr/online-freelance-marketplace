<script lang="ts">
    import LeftSignSideLayout from "@/common-components/Sign/components/LeftSignSide/LeftSignSideLayout.svelte";
    import LeftSignSideTopText
        from "@/common-components/Sign/components/LeftSignSideTopText/LeftSignSideTopText.svelte";
    import SignForm from "@/common-components/Sign/components/SignForm/SignForm.svelte";
    import SubFormText from "@/common-components/Sign/components/SubFormText/SubFormText.svelte";

    const signProps = $props();
</script>

<LeftSignSideLayout>
    <LeftSignSideTopText {...signProps}/>
    <SignForm {...signProps}/>
    <SubFormText {...signProps}/>
</LeftSignSideLayout>
