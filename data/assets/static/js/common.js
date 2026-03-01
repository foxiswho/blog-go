/**
 * 将form表单内的元素序列化为对象
 * @param form
 * @returns {{}}
 */
function serializeObject(form) {
    var o = {};
    $.each($(form).serializeArray(), function (index) {
        if (o[this['name']]) {
            o[this['name']] = o[this['name']] + "," + this['value'];
        } else {
            o[this['name']] = this['value'];
        }
    });
    return o;
}
;(function ($, window, document, undefined) {
    $.fn.value = function (options) {
        var _selector = this.selector, $this = $(_selector), val;
        if ($this.length <= 0) {
            var first = _selector.substr(0, 1);
            if ("#" === first || "." === first) {
                $this = $(_selector);
            } else {
                $this = $("[name='" + _selector + "']");
            }
            if ($this.length <= 0) {
                console.info(_selector + ' 不存在 ');
                return false;
            }
        }
        if (options === undefined) {
            if ($this.eq(0).is(":radio")) { //单选按钮
                val = $this.filter(":checked").val();
            } else if ($this.eq(0).is(":checkbox")) { //复选框
                val = '';
                $this.filter(":checked").each(function (i) {
                    val += (i == 0 ? '' : ',') + $(this).val()
                });
            } else {
                val = $this.val();
            }
            //判断是否是数值文本框
            if ($this.attr('type') == 'number') {
                if (isNaN(val)) {
                    val = 0;
                } else if (val == '') {
                    val = 0;
                }
            }
        } else {
            //判断是否是数值文本框
            if ($this.eq(0).is(":radio")) {
                $this.filter("[value='" + options + "']").each(function () {
                    this.checked = true
                });
                return true;
            } else if ($this.eq(0).is(":checkbox")) {
                if (!$.isArray(options) && options && options.indexOf(',') > 0) {
                    $this.val(options.split(','));
                } else if (options == '' || options === false) {
                    $this.attr('checked', false);
                }
                return true;
            } else {
                $this.val(options);
            }
            return true;
        }
        return val;
    }
})(jQuery, window, document);