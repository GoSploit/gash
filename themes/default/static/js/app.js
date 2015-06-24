$(function(){
	main = {
		projects: ko.observableArray()
	};
	main.projects.push({
		attackType: ko.observable("straight"),
		hashes: ko.observable(6)
	});
	
	
	
	ko.applyBindings(main);
	
	
	
	
	
})