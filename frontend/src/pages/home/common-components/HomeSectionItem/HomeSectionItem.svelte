<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import HomeSectionItemIcon from "@/pages/home/common-components/HomeSectionItemIcon/HomeSectionItemIcon.svelte";
    import HomeSectionItemTitle from "@/pages/home/common-components/HomeSectionItemTitle/HomeSectionItemTitle.svelte";
    import HomeSectionItemDescription
        from "@/pages/home/common-components/HomeSectionItemDescription/HomeSectionItemDescription.svelte";

    type HowItWorkSectionStepProps = {
        icon: string;
        title: string;
        description: string;
    };

    const props: HowItWorkSectionStepProps = $props();
</script>

<PaperElement styles="flex gap-2 flex-col justify-center items-center xl:aspect-square">
    <HomeSectionItemIcon {...props}/>
    <HomeSectionItemTitle {...props}/>
    <HomeSectionItemDescription {...props}/>
</PaperElement>