<template>
	<div>
		<b-container fluid class="text-light text-center">
			<b-row>
				<b-col
					cols="4"
					class="btn-primary"
					:class="[commonClass]"
				>
					assemble
				</b-col>
				<b-col
					cols="4"
					class="btn-primary"
					:class="[commonClass]"
				>
					complete
				</b-col>
				<b-col
					cols="4"
					class="btn-primary"
					:class="[commonClass]"
				>
					passed
				</b-col>
			</b-row>
		</b-container>

		<b-container fluid class="text-light text-center">
			<b-row>
				<b-col
					cols="4"
				>
					<b-row>
						<b-col cols="12"
						class="btn-info"
						:class="[commonClass]">
							a
						</b-col>
					</b-row>
				</b-col>

				<b-col
					cols="4"
				>
					<b-row>
						<b-col cols="12"
						class="btn-info"
						:class="[commonClass]">
							b
						</b-col>
					</b-row>
				</b-col>

				<b-col
					cols="4"
				>
					<b-row>
						<b-col cols="12"
						class="btn-info"
						:class="[commonClass]">
							c
						</b-col>
					</b-row>
				</b-col>

			</b-row>
		</b-container>
	</div>
</template>

<script>
module.exports = {
	props: {
		'commonClass': {type: String, default: ''},
	},

	methods: {
	},

	components: {
	},


};
</script>
