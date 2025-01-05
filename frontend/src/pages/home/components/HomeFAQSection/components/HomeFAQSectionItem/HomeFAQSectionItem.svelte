<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import HomeFAQSectionItemQuestion
        from "@/pages/home/components/HomeFAQSection/components/HomeFAQSectionItemQuestion/HomeFAQSectionItemQuestion.svelte";
    import HomeFAQSectionItemAnswer
        from "@/pages/home/components/HomeFAQSection/components/HomeFAQSectionItemAnswer/HomeFAQSectionItemAnswer.svelte";

    type HomeFAQSectionItemProps = {
        question: string,
        answer: string
    };

    const props: HomeFAQSectionItemProps = $props();
</script>

<PaperElement styles="w-full text-start">
    <HomeFAQSectionItemQuestion {...props}/>
    <HomeFAQSectionItemAnswer {...props}/>
</PaperElement>
