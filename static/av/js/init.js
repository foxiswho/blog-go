eva.ready(function(){
	if($.browser.msie && (($.browser.version == "6.0" && !$.support.style) || $.browser.version == "7.0")) {
		jQuery.fx.off = true;
	}

	$("a[rel=external]").live("click", function(){
		$(this).attr("target", "_blank");		
	});

	$("span[data-url-wrap]").each(function(){
		$(this).wrapInner('<a href="' + $(this).attr('data-url-wrap') + '"></a>');
	});

	eva.loader('/static/av/js/addthis_widget.js#pubid=allovince', function(){});
	//
	$(".agotime").each(function(){
		var time = $(this);
		time.html(moment(time.attr("datetime"), "YYYY-MM-DDTHH:mm:ss ZZ").fromNow());
	});

	if($("#blog pre>code")[0]){
		eva.loadcss(eva.assets(['/static/av/css/github.css']));
		eva.loader(eva.assets(['/static/av/js/highlight.min.js']), function(){
			hljs.tabReplace = '    ';
			hljs.initHighlightingOnLoad();
		});
	}

	//TODO: 检查至少有两级标题
	var titles = $("#blog > article :header");
	// eva.p($.unique(titles.toArray()));
	if(titles.length > 4){
		eva.loader(eva.assets(['/static/av/js/jquery.toc.min.js', '/static/av/js/jquery.jscrollpane.min.js']), function(){
			var toc = $('<div id="toc" class="typo typocn min"></div>').appendTo('body');
			toc.toc({
				'selectors': 'h1,h2,h3,h4', //elements to use as headings
				'container': '#blog', //element to find all selectors in
				'onHighlight': function(el) {
					var offset = $(el).position();
					toc.data('jsp').scrollTo(0, offset.top - 200);
				},
				'smoothScrolling': true
			});
			toc.jScrollPane({
				//animateScroll: true
			});	
			toc.append('<span class="handler"><i class="icon-angle-left"></i></span>');

			var handler = toc.find('.handler');
			var openToc = function(){
				if(toc.is(':animated')){
					return false;
				}
				if(toc.hasClass('min')){
					handler.hide();
					toc.animate({right:0}, "fast", function(){
						toc.toggleClass('min');
						handler.find('i').toggleClass('icon-angle-left icon-angle-right');
						handler.show();
					});				
				}		
			}
			var closeToc = function(){
				if(toc.is(':animated')){
					return false;
				}
				if(!toc.hasClass('min')){
					handler.hide();
					toc.animate({right:'-300px'}, "fast", function(){
						toc.toggleClass('min');
						handler.find('i').toggleClass('icon-angle-left icon-angle-right');
						handler.show();
					});				
				}
			}
			handler.on('click', function(){
				if(toc.is(':animated')){
					return false;
				}
				if(toc.hasClass('min')){
					openToc();
				} else {
					closeToc();
				}
			});

			//show by default
			if($(document).width() - 1200 > 300){
				openToc();
			}
		});

	}

	eva.story();

	if($("#timeline-embed")[0]){
		var $timeline = $("#timeline-embed");
		$("body").addClass('timeline');
		eva.loader(eva.assets(['/static/av/js/storyjs-embed.js']), function(){
			var minHeight = 550;
			var height = $(window).height() > minHeight ? $(window).height() - 85 : minHeight;
			var source = $timeline.html();
			//source = eva.d('proxy/') + '?page=' + encodeURIComponent(source.replace(/&amp;/g, '&'));
			source = eva.d('/proxy/spreadsheet');
			$timeline.empty();
			var config = {
				type:               'timeline',
				width:              '100%',
				height:             height,
				start_at_end:       true, 
				debug:              false,
				lang:               'zh-cn',
				source:             source,
				embed_id:           'timeline-embed',
				js:                 eva.assets('/static/av/js/timeline.min.js'),
				css:                eva.assets('/static/av/css/timeline.css')
			};
			//console.log(config);
			createStoryJS(config);
		});	
	}
    //百度统计
    eva.loader('https://hm.baidu.com/hm.js?1e3a014af6088caa70251a766ac9e6d4', function(){});
    //google 翻译
    eva.loader('http://translate.google.com/translate_a/element.js?cb=googleTranslateElementInit', function(){});
});

eva.destruct = function(){
	//eva.login.init();
	//$.evaValidate.init();
};

eva.story = function(){

	if(!$("#stories")[0]){
		return false;
	}

	var startStory = function(){
		var maxLoaded = 10;
		var loadTimes = 1;
		var container = $("#stories>.stories-wrap>.stories");
		var loader = $("#stories .load-story");
		$("body").append('<div id="load-area"></div>');
		var loadArea = $("#load-area");
		var loaded = [];

		var initStory = function(items){

			items.each(function(){
					
				var item = $(this);
				item.find(".agotime").each(function(){
					var time = $(this);
					time.html(moment(time.attr("datetime"), "YYYY-MM-DDTHH:mm:ss ZZ").fromNow());
				});
				if(item.hasClass("inited")){
					return true;
				}
				if(item.hasClass("type-note")){
					item.find(">.content").jScrollPane({autoReinitialise: true});	
				}

				item.addClass("inited");

				return true;
			});

		};

		$(window).resize(function(){
			container.masonry({
				itemSelector : '.item',
				columnWidth : $(window).width() > 800 ? 320 : 260,
				isAnimated: true
			}).masonry( 'reload' );
		});

		container.imagesLoaded( function(){
			container.masonry({
				itemSelector : '.item',
				columnWidth : $(window).width() > 800 ? 320 : 260,
				isAnimated: true
			});

			var items = container.find(".item:not(.inited)");
			//eva.p(items.length);
			initStory(items);
		});

		loadArea.hide();
		function inArray(stringToSearch, arrayToSearch) {
			for (s = 0; s < arrayToSearch.length; s++) {
				thisEntry = arrayToSearch[s].toString();
				if (thisEntry == stringToSearch) {
					return true;
				}
			}
			return false;
		}
		
		var loadNewStory = function(loader){

			if(loadTimes > maxLoaded) {
				return false;
			}

			var url = loader.attr("href");
			if(inArray(url, loaded)){
				return false;
			}

			loader.addClass("disabled").html(" （；^ω^） 正在努力载入...");
			loaded.push(url);

			loadArea.load(url + ' #stories>.stories-wrap', function() {
				var newUrl = loadArea.find(".load-story").attr("href");
				loader.attr("href", newUrl); 
				var content = loadArea.find(".stories-wrap>.stories>.item");
				loadArea.imagesLoaded( function(){
					container.append(content).masonry( 'appended', content, true);
					initStory(container.find(".item:not(.inited)"));
					loadArea.html('');
					loader.removeClass("disabled").html("更多");
				});
				loadTimes++;
			});

			return false;
		};

		$(window).scroll(function () { 
			var pageH = $(document).height(); //页面总高度 
			var scrollT = $(window).scrollTop(); //滚动条top 
			var winH = $(window).height(); 
			var offset = pageH - scrollT - winH;
			if(offset < 300){
				loadNewStory(loader);
			}
		}); 
	}

	eva.loader(eva.s([
		'/static/av/js/jquery.masonry.js',
		'/static/av/js/jquery.mousewheel.js',
		'/static/av/js/jquery.jscrollpane.js'
	]), startStory);
};

eva.init(eva_config);
