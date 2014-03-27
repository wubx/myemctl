<!DOCTYPE html>

<html>
<head>
<title>天眼v0.01</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<script src="static/js/jquery.min.js"></script>
<script src="static/js/highstock.js"></script>
<link rel="stylesheet" href="static/css/bootstrap.min.css">

</head>
  	
<body>
<div class="row">
	<div class="col-md-1"> </div>
	<div class="col-md-10 jumbotron container">
	<h1> 天眼监控 </h1>
	</div>
</div>

<div class="row">
	<div class="col-md-1"></div>
	<div class="col-md-10" id="cpu"></div>
	
</div>

<div class="row">
	<div class="col-md-1"></div>
	<div class="col-md-10" id="load"></div>
	
</div>
<div class="row">
	<div class="col-md-1"></div>
	<div class="col-md-10" id="mem"></div>
	
</div>

<div class="row">
	<div class="col-md-1"></div>
	<div class="col-md-10" id="bphit"></div>
	
</div>

<div class="row">
	<div class="col-md-1"></div>
	<div class="col-md-10" id="myactive"></div>
	
</div>

<div class="row">
	<div class="col-md-1"></div>
	<div class="col-md-10" id="myconn"></div>
	
</div>





<div class="row">
	<div class="col-md-12 well container">

			    <p class="description container">
				天眼监控用MySQL运行情况分析,类似于Oracle AWR性能数据收集及分析
			    <br />
			    	Official website: <a href="http://{{.Website}}">{{.Website}}</a>
			    <br />
			    	Contact me: {{.Email}}
			    </p>
	</div>
</div>


<script>
// 用于CPU画图展示
		$(function() {

				$.getJSON('/cpu', function(data) {
					//console.debug(data);
					$('#cpu').highcharts('StockChart', {


					rangeSelector : {
					selected : 2
					},

					title : {
					text : 'CPU性能展示'
					},
				     	legend: {
						layout: 'vertical',
						align: 'right',
						verticalAlign: 'middle',
						borderWidth: 0
					    },

					series : [
						{
						name : 'CpuIdle',
						data : data.CpuIdle,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'CpuTotal',
						data : data.CpuTotal,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'CpuUser',
						data : data.CpuUser,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'CpuSys',
						data : data.CpuSys,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'CpuIowait',
						data : data.CpuIowait,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'CpuIrq',
						data : data.CpuIrq,
						tooltip: {
							valueDecimals: 2
							}
						}
					],
					});
				});

			});
</script>

<script>
//load 性能展示
		$(function() {

				$.getJSON('/load', function(data) {
					//console.debug(data);
					$('#load').highcharts('StockChart', {


					rangeSelector : {
					selected : 2
					},

					title : {
					text : '系统Load性能展示'
					},
				     	legend: {
						layout: 'vertical',
						align: 'right',
						verticalAlign: 'middle',
						borderWidth: 0
					    },

					series : [
						{
						name : 'La1',
						data : data.La1,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'La5',
						data : data.La5,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'La15',
						data : data.La15,
						tooltip: {
							valueDecimals: 2
							}
						}
					],
					});
				});

			});
</script>

<script>
//mem使用情况
		$(function() {

				$.getJSON('/mem', function(data) {
					//console.debug(data);
					$('#mem').highcharts('StockChart', {


					rangeSelector : {
					selected : 2
					},

					title : {
					text : '系统Mem使用展示(单位M)'
					},
				     	legend: {
						layout: 'vertical',
						align: 'right',
						verticalAlign: 'middle',
						borderWidth: 0
					    },

					series : [
						{
						name : 'MemUse',
						data : data.Use,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'MemFree',
						data : data.Free,
						tooltip: {
							valueDecimals: 2
							}
						},
						{
						name : 'MemTotal',
						data : data.Total,
						tooltip: {
							valueDecimals: 2
							}
						}
						],
					});
				});

			});

</script>

<script>
//mythd
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

<script>
//Mysql Innodb bp hit
               $(function() {

                                $.getJSON('/bphit', function(data) {
                                        //console.debug(data);
                                        $('#bphit').highcharts('StockChart', {


                                        rangeSelector : {
                                        selected : 2
                                        },

                                        title : {
                                        text : 'MySQL Innodb bp hit'
                                        },
                                        legend: {
                                                layout: 'vertical',
                                                align: 'right',
                                                verticalAlign: 'middle',
                                                borderWidth: 0
                                            },

                                        series : [
                                                {
                                                name : 'BpHit',
                                                data : data.Hit,
                                                tooltip: {
                                                        valueDecimals: 2
                                                        }
                                                }
                                            ],
                                        });
                                });

                        });
</script>

<script>
//MySQL active 展示
               $(function() {

                                $.getJSON('/active', function(data) {
                                        //console.debug(data);
                                        $('#myactive').highcharts('StockChart', {


                                        rangeSelector : {
                                        selected : 2
                                        },

                                        title : {
                                        text : 'MySQL活跃情况'
                                        },
                                        legend: {
                                                layout: 'vertical',
                                                align: 'right',
                                                verticalAlign: 'middle',
                                                borderWidth: 0
                                            },

                                        series : [
                                                {
                                                name : 'Select',
                                                data : data.Sselect,
                                                tooltip: {
                                                        valueDecimals: 2
                                                        }
                                                },
                                                {
                                                name : 'Insert',
                                                data : data.Iinsert,
                                                tooltip: {
                                                        valueDecimals: 2
                                                        }
                                                },
                                                {
                                                name : 'Update',
                                                data : data.Uupdate,
                                                tooltip: {
                                                        valueDecimals: 2
                                                        }
                                                },
                                                {
                                                name : 'Delete',
                                                data : data.Ddelete,
                                                tooltip: {
                                                        valueDecimals: 2
                                                        }
                                                },
                                                {
                                                name : 'Call',
                                                data : data.Ccall,
                                                tooltip: {
                                                        valueDecimals: 2
                                                        }
                                                }
                                        ],
                                        });
                                });

                        });
</script>
</body>
</html>
