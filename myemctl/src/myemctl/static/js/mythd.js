<script>
// 用于CPU画图展示
		$(function() {

				$.getJSON('/thd', function(data) {
					//console.debug(data);
					$('#myconn').highcharts('StockChart', {


					rangeSelector : {
					selected : 2
					},

					title : {
					text : 'MySQL连接数'
					},
				     	legend: {
						layout: 'vertical',
						align: 'right',
						verticalAlign: 'middle',
						borderWidth: 0
					    },

					series : [
						{
						name : 'ThdConned',
						data : data.ThdConn,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'ThdRun',
						data : data.ThdRun,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'ThdCached',
						data : data.ThdCache,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'ThdAborted',
						data : data.ThdAborted,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'ThdCreated',
						data : data.ThdCreated,
						tooltip: {
							valueDecimals: 2
							}
						}
					],
					});
				});

			});
</script>
