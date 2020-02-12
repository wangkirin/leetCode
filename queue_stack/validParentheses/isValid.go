package validParentheses

//No. 0020 Valid Parentheses

func isValid(s string) bool {

	paranthese:=[]byte(s)
	stack:=[]byte{}
	lenParathese:=len(paranthese)
	if lenParathese%2!=0 {
		return false
	}
	for i:=(lenParathese-1);i>=0 ;i--  {
		if paranthese[i]=='}'||paranthese[i]==']'||paranthese[i]==')' {
			stack=append(stack,paranthese[i])
		}
		if paranthese[i]=='{'{
			if len(stack)>0&&stack[len(stack)-1]=='}'{
				stack=stack[:len(stack)-1]
		}else {
				return false
			}
		}
		if paranthese[i]=='['{
			if len(stack)>0&&stack[len(stack)-1]==']'{
				stack=stack[:len(stack)-1]
			}else {
				return false
			}
		}
		if paranthese[i]=='('{
			if len(stack)>0&&stack[len(stack)-1]==')'{
				stack=stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}
	if len(stack)>0 {
		return false
	}
	return true
}