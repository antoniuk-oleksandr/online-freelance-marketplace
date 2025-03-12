<script lang="ts">
    type TitleProps = {
        text: string,
        capitalize?: boolean,
    };

    const {text, capitalize}: TitleProps = $props();
</script>

<h2 class="text-xl {capitalize === false ? '' : 'capitalize'} font-semibold">{text}</h2>
