<template>
	<b-container fluid class="text-light text-center">
		<b-row>
			<b-col
				cols="6"
				:class="[commonClass, classSelected('ambient')]"
				@click="methodChangeDevice('ambient')"
			>
				<icon-device-ambient />
			</b-col>

			<b-col
				cols="6"
				:class="[commonClass, classSelected('projector')]"
				@click="methodChangeDevice('projector')"
			>
				<icon-device-projector />
			</b-col>
		</b-row>

		<br />

		<div v-show="device=='ambient'">
			<device-ambient
				:common-class="commonClass"
				:light-status="lightStatus"
				:aircon-status="airconStatus"
				:aircon-warm="airconWarm"
				:aircon-cool="airconCool"
				@send-command="methodSendCommand"
			/>
		</div>

		<div v-show="device=='projector'">
			<device-projector
				:common-class="commonClass"

				@send-command="methodSendCommand"
			/>
		</div>
	</b-container>
</template>

<script>
module.exports = {
	components: {
		'icon-device-light': httpVueLoader('web/vue/icon/device-light.vue'),
		'icon-device-aircon': httpVueLoader('web/vue/icon/device-aircon.vue'),
		'icon-device-ambient': httpVueLoader('web/vue/icon/device-ambient.vue'),
		'icon-device-projector': httpVueLoader('web/vue/icon/device-projector.vue'),

		'device-light': httpVueLoader('web/vue/device/light.vue'),
		'device-aircon': httpVueLoader('web/vue/device/aircon.vue'),
		'device-ambient': httpVueLoader('web/vue/device/ambient.vue'),
		'device-projector': httpVueLoader('web/vue/device/projector.vue'),

},
	props: {
		'commonClass': {type: String, default: ''},
		'device': {type: String, default: ''},
		'lightStatus': {type: String, default: ''},
		'airconStatus': {type: String, default: ''},
		'airconWarm': {type: Number , default: 0},
		'airconCool': {type: Number , default: 0},
	},


	methods: {
		classSelected (value) {
			return value == this.device ? 'btn-primary' : 'btn-secondary';
		},

		methodChangeDevice: function(value) {
			this.$emit('change-device', value);
		},

		methodSendCommand: function(value) {
			this.$emit('send-command', value);
		},
	},
}


