export namespace backend {
	
	export class Config {
	    Id: number;
	    Name: string;
	    Value: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Name = source["Name"];
	        this.Value = source["Value"];
	    }
	}
	export class Conn {
	    Id: number;
	    Name: string;
	    URL: string;
	    Token: string;
	
	    static createFrom(source: any = {}) {
	        return new Conn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Name = source["Name"];
	        this.URL = source["URL"];
	        this.Token = source["Token"];
	    }
	}
	export class LogMsg {
	    Time: number;
	    Level: string;
	    Msg: string;
	
	    static createFrom(source: any = {}) {
	        return new LogMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Time = source["Time"];
	        this.Level = source["Level"];
	        this.Msg = source["Msg"];
	    }
	}
	export class Msg {
	    Id: number;
	    TopicId: number;
	    Content: string;
	    Type: string;
	    Time: number;
	
	    static createFrom(source: any = {}) {
	        return new Msg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.TopicId = source["TopicId"];
	        this.Content = source["Content"];
	        this.Type = source["Type"];
	        this.Time = source["Time"];
	    }
	}
	export class Topic {
	    Id: number;
	    ConnID: number;
	    Name: string;
	    Topic: string;
	
	    static createFrom(source: any = {}) {
	        return new Topic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.ConnID = source["ConnID"];
	        this.Name = source["Name"];
	        this.Topic = source["Topic"];
	    }
	}

}

