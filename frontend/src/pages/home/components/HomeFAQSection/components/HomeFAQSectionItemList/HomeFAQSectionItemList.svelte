<script lang="ts">
    import HomeFAQSectionItemListLayout
        from "@/pages/home/components/HomeFAQSection/components/HomeFAQSectionItemList/HomeFAQSectionItemListLayout.svelte";
    import HomeFAQSectionItem
        from "@/pages/home/components/HomeFAQSection/components/HomeFAQSectionItem/HomeFAQSectionItem.svelte";

    const faqs = [
        {
            question: "How do I hire a freelancer?",
            answer: "You can use our search page to find freelancers, place an order-request with secure payment, and start collaborating."
        },
        {
            question: "What payment methods are accepted?",
            answer: "We accept all major credit cards and other secure payment methods."
        },
        {
            question: "Can I cancel an order-request?",
            answer: "Yes, you can cancel an order-request before it is marked as completed by the freelancer."
        },
        {
            question: "Are freelancers verified?",
            answer: "Yes, all freelancers go through a verification process to ensure quality and authenticity."
        }
    ];
</script>

<HomeFAQSectionItemListLayout>
    {#each faqs as faq}
        <HomeFAQSectionItem {...faq}/>
    {/each}
</HomeFAQSectionItemListLayout>