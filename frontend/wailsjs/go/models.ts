export namespace main {
	
	export class Task {
	    id: number;
	    title: string;
	    status: boolean;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.status = source["status"];
	        this.created_at = source["created_at"];
	    }
	}

}

