<script>
    import FooterIconLayout from "@/common-components/Footer/components/FooterIcon/FooterIconLayout.svelte";
    import Icon from "@iconify/svelte";

    export let link;
    export let iconName;
</script>

<FooterIconLayout link={link}>
    <Icon icon={iconName} />
</FooterIconLayout>