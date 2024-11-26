<script lang="ts">
    import SuccessMessageLayout from "@/common-components/SuccessMessage/SuccessMessageLayout.svelte";
    import SuccessMessageTitle
        from "@/common-components/SuccessMessage/components/SuccessMessageTitle/SuccessMessageTitle.svelte";
    import SuccessMessageDescription
        from "@/common-components/SuccessMessage/components/SuccessMessageDescription/SuccessMessageDescription.svelte";
    import SuccessMessageSubtext
        from "@/common-components/SuccessMessage/components/SuccessMessageSubtext/SuccessMessageSubtext.svelte";
    import SuccessMessageHomeButton
        from "@/common-components/SuccessMessage/components/SuccessMessageHomeButton/SuccessMessageHomeButton.svelte";
    import SuccessMessageIcon
        from "@/common-components/SuccessMessage/components/SuccessMessageIcon/SuccessMessageIcon.svelte";

    type SuccessMessageProps = {
        title: string,
        description: string,
        subText: string,
    }

    const {title, description, subText}: SuccessMessageProps = $props();
</script>

<SuccessMessageLayout>
    <SuccessMessageIcon/>
    <SuccessMessageTitle {title}/>
    <SuccessMessageDescription {description}/>
    <SuccessMessageSubtext {subText}/>
    <SuccessMessageHomeButton/>
</SuccessMessageLayout>
