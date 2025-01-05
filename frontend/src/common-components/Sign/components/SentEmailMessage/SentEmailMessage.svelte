<script>
    import Icon from "@iconify/svelte";
    import SentEmailMessageLayout
        from "@/common-components/Sign/components/SentEmailMessage/SentEmailMessageLayout.svelte";
</script>

<SentEmailMessageLayout>
    <Icon icon="material-symbols:check-circle" class="text-green-500" width="64" height="64"/>
    <h1 class="text-2xl font-bold text-center">Email Was Sent Successfully</h1>
    <p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
        We've sent a confirmation email to your inbox.
        Please check your email and click the confirmation link
        to proceed.
    </p>
    <p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
        If you don’t see the email, check your spam folder.
    </p>
</SentEmailMessageLayout>
